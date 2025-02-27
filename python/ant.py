import pygame
import sys
import time

# --------------------------------------------
# Settings --------------------------------
# --------------------------------------------
WINDOW_WIDTH, WINDOW_HEIGHT = 800, 600
ROOM_SIZE        = 40  # regular room size
BIG_ROOM_SIZE    = 60  # start / end room size
ANT_SIZE         = 10
ANIMATION_STEPS  = 30  # frames per movement step
FPS              = 30 # frames per second

# colors
WHITE       = (139, 69, 19)
START_COLOR = (0, 200, 0)
END_COLOR   = (200, 0, 0)
ROOM_COLOR  = (255, 215, 0)
LINE_COLOR  = (255, 255, 255)
ANT_COLOR   = (200, 200, 200)

pygame.init()
screen = pygame.display.set_mode((WINDOW_WIDTH, WINDOW_HEIGHT))
pygame.display.set_caption("Ant Movement - Step/Backward/Auto")
clock = pygame.time.Clock()

# --------------------------------------------------
# 1) read file output.txt
# --------------------------------------------------
def read_file(file_path):
    """
    reads the file and returns:
      - num_ants        (int)
      - rooms           ({room_name: (x, y)})
      - start_room      (str)
      - end_room        (str)
      - paths           (list of tuples [(r1, r2), ...])
      - moves           (list of [ (ant_id, room_name), ... ] ανά γύρο)
    """
    with open(file_path, 'r') as f:
        lines = [line.strip() for line in f if line.strip()]

    num_ants = int(lines[0])
    rooms = {}
    paths = []
    moves = []
    start_room = None
    end_room = None

    i = 1

    def parse_room_line(line):
        # for example "0 1 4" -> (name="0", x=1, y=4)
        parts = line.split()
        return str(parts[0]), int(parts[1]), int(parts[2])

    while i < len(lines):
        line = lines[i].strip()
        if line.lower().startswith("total turns"):
            # end of file
            break

        if line.startswith('L'):
            # movements turn for example "L1-2 L2-4"
            round_moves = []
            tokens = line.split()
            for tk in tokens:
                left, roomname = tk.split('-')  # for example left="L1", roomname="2"
                ant_id = int(left[1:])          # remove 'L'
                round_moves.append((ant_id, roomname))
            moves.append(round_moves)
            i += 1
            continue

        if '-' in line and not line.startswith('L'):
            # room connections for example "0-1"
            r1, r2 = line.split('-')
            paths.append((r1, r2))
            i += 1
            continue

        if line == "##start":
            # next line is start room
            sline = lines[i+1]
            nm, x, y = parse_room_line(sline)
            rooms[nm] = (x,y)
            start_room = nm
            i += 2
            continue

        if line == "##end":
            # next line is end room
            eline = lines[i+1]
            nm, x, y = parse_room_line(eline)
            rooms[nm] = (x,y)
            end_room = nm
            i += 2
            continue
        if line.startswith("#") and not line.startswith("##"):
            print('here')
            i += 1
            continue

        # regular room
        nm, x, y = parse_room_line(line)
        rooms[nm] = (x,y)
        i += 1

    return num_ants, rooms, start_room, end_room, paths, moves

# --------------------------------------------------
# 2) Build round positions
# --------------------------------------------------
def build_round_positions(num_ants, start_room, moves):

    # all ants start from the start_room
    current = [start_room]*num_ants
    round_positions = [current[:]]  # copy list

    for round_moves in moves:
        # copy current positions to edit them
        new_pos = current[:]
        # apply movements
        for (ant_id, roomname) in round_moves:
            new_pos[ant_id-1] = roomname
        # save positions
        round_positions.append(new_pos)
        # update current positions
        current = new_pos
    return round_positions

# --------------------------------------------------
# 3) Design colony
# --------------------------------------------------
def get_room_rect(room_name, rooms, start_room, end_room):
    
    (rx, ry) = rooms[room_name]
    if room_name == start_room or room_name == end_room:
        return (rx*ROOM_SIZE, ry*ROOM_SIZE, BIG_ROOM_SIZE, BIG_ROOM_SIZE)
    else:
        return (rx*ROOM_SIZE, ry*ROOM_SIZE, ROOM_SIZE, ROOM_SIZE)

def get_room_center(room_name, rooms, start_room, end_room):
   
    (rx, ry) = rooms[room_name]
    if room_name == start_room or room_name == end_room:
        # Κέντρο του BIG_ROOM_SIZE
        return (rx*ROOM_SIZE + BIG_ROOM_SIZE//2,
                ry*ROOM_SIZE + BIG_ROOM_SIZE//2)
    else:
        return (rx*ROOM_SIZE + ROOM_SIZE//2,
                ry*ROOM_SIZE + ROOM_SIZE//2)

def draw_rooms(rooms, start_room, end_room):
    
    for rname in rooms:
        rect = get_room_rect(rname, rooms, start_room, end_room)
        if rname == start_room:
            color = START_COLOR
        elif rname == end_room:
            color = END_COLOR
        else:
            color = ROOM_COLOR
        pygame.draw.rect(screen, color, rect)
        pygame.draw.rect(screen, LINE_COLOR, rect, 2)
        # room name
        font = pygame.font.SysFont(None, 18)
        text = font.render(rname, True, (0,0,0))
        screen.blit(text, (rect[0]+3, rect[1]+3))

def draw_lines(paths, rooms, start_room, end_room):
    """design connection lines"""
    for (r1, r2) in paths:
        if r1 not in rooms or r2 not in rooms:
            continue
        c1 = get_room_center(r1, rooms, start_room, end_room)
        c2 = get_room_center(r2, rooms, start_room, end_room)
        pygame.draw.line(screen, LINE_COLOR, c1, c2, 3)

def draw_ants(positions, rooms, start_room, end_room):
    """
   Design ants
    """
    for i, room_name in enumerate(positions):
        if room_name not in rooms:
            continue
        (cx, cy) = get_room_center(room_name, rooms, start_room, end_room)
        pygame.draw.circle(screen, ANT_COLOR, (cx, cy), ANT_SIZE)
        font = pygame.font.SysFont(None, 16)
        text = font.render(str(i+1), True, (0,0,0))
        screen.blit(text, (cx - 5, cy - 5))

# --------------------------------------------------
# 4) movements animation
# --------------------------------------------------
def animate_transition(rooms, start_room, end_room,
                       from_positions, to_positions):
    """
    Δημιουργεί ομαλή μετάβαση (ANIMATION_STEPS καρέ) μεταξύ
    των θέσεων from_positions και to_positions.
    from_positions / to_positions: λίστες με ονόματα δωματίων.
    """
    for frame in range(ANIMATION_STEPS):
       
        t = frame / (ANIMATION_STEPS - 1)
        
        
        screen.fill(WHITE)
        draw_rooms(rooms, start_room, end_room)
        draw_lines(paths, rooms, start_room, end_room)

        
        for i in range(len(from_positions)):
            rname1 = from_positions[i]
            rname2 = to_positions[i]
            x1, y1 = get_room_center(rname1, rooms, start_room, end_room)
            x2, y2 = get_room_center(rname2, rooms, start_room, end_room)
            # interpolation
            cx = x1 + (x2 - x1)*t
            cy = y1 + (y2 - y1)*t
            
            pygame.draw.circle(screen, ANT_COLOR, (cx, cy), ANT_SIZE)
            font = pygame.font.SysFont(None, 16)
            text = font.render(str(i+1), True, (0,0,0))
            screen.blit(text, (cx - 5, cy - 5))

        pygame.display.flip()
       
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                pygame.quit()
                sys.exit()
        clock.tick(FPS)

# --------------------------------------------------
# 5) game loop
# --------------------------------------------------
def game_loop(file_path):
    global paths  
    num_ants, rooms, start_room, end_room, paths_, moves = read_file(file_path)
    paths = paths_  # αποθηκεύουμε σε global για use σε animate_transition

    
    round_positions = build_round_positions(num_ants, start_room, moves)
    total_rounds = len(round_positions) - 1  

    current_round = 0  # start from turn 0
    done = False

    while not done:
       
        screen.fill(WHITE)
        draw_rooms(rooms, start_room, end_room)
        draw_lines(paths, rooms, start_room, end_room)
        draw_ants(round_positions[current_round], rooms, start_room, end_room)
        pygame.display.flip()

        
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                done = True

            elif event.type == pygame.KEYDOWN:
                # space key to play all movements
                if event.key == pygame.K_SPACE:
                   
                    while current_round < total_rounds:
                        animate_transition(
                            rooms, start_room, end_room,
                            round_positions[current_round],
                            round_positions[current_round+1]
                        )
                        current_round += 1
                        screen.fill(WHITE)
                        draw_rooms(rooms, start_room, end_room)
                        draw_lines(paths, rooms, start_room, end_room)
                        draw_ants(round_positions[current_round], rooms, start_room, end_room)
                        pygame.display.flip()
                        # μικρή παύση  (προαιρετικά)
                        time.sleep(0.1)
                        
                # right or down key for next movement
                elif event.key in (pygame.K_RIGHT, pygame.K_DOWN):
                    if current_round < total_rounds:
                        animate_transition(
                            rooms, start_room, end_room,
                            round_positions[current_round],
                            round_positions[current_round+1]
                        )
                        current_round += 1

                # left or right key for previous movement
                elif event.key in (pygame.K_LEFT, pygame.K_UP):
                    if current_round > 0:
                        animate_transition(
                            rooms, start_room, end_room,
                            round_positions[current_round],
                            round_positions[current_round-1]
                        )
                        current_round -= 1

        clock.tick(FPS)

    pygame.quit()
    sys.exit()

# --------------------------------------------------
if __name__ == "__main__":
    game_loop("output.txt")
