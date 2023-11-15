package main

type level struct {
	ghostMap       map[position]*ghost
	tombstoneMap   map[position]*tombstone
	width          int
	height         int
	door           position
	playerStartPos position
}

func makeLevels(isDemo bool) []level {
	if isDemo {
		return []level{
			makeLevel(demoTombstoneLevel()),
			makeLevel(demoWanderLevel()),
			makeLevel(demoFollowLevel()),
			makeLevel(demoHuntLevel()),
		}
	}

	return []level{
		makeLevel(level0()),
		makeLevel(level1()),
		makeLevel(level2()),
		makeLevel(level3()),
		makeLevel(level4()),
		makeLevel(level5()),
		makeLevel(level6()),
		makeLevel(level7()),
		makeLevel(level8()),
		makeLevel(level9()),
	}
}

func makeLevel(data []string) level {
	lvl := level{
		width:        len(data[0]),
		height:       len(data),
		tombstoneMap: map[position]*tombstone{},
		ghostMap:     map[position]*ghost{},
	}

	for rowIdx, row := range data {
		for colIdx, val := range row {
			mapPoint := position{y: rowIdx, x: colIdx}

			switch val {
			case 'P':
				lvl.playerStartPos = position{x: colIdx, y: rowIdx}
			case 'S':
				lvl.tombstoneMap[mapPoint] = &tombstone{}
			case 'K':
				lvl.tombstoneMap[mapPoint] = &tombstone{hasKey: true}
			case 'D':
				lvl.door = position{x: colIdx, y: rowIdx}
			case 'W':
				lvl.ghostMap[mapPoint] = &ghost{kind: "wander"}
			case 'F':
				lvl.ghostMap[mapPoint] = &ghost{kind: "follow"}
			case 'H':
				lvl.ghostMap[mapPoint] = &ghost{kind: "hunt"}
			}
		}
	}

	return lvl
}

func level0() []string {
	return []string{
		"                    D                           ",
		"                                                ",
		"                                S               ",
		"                                         S      ",
		"          S        S      S                     ",
		"          S                        S            ",
		"          S                                     ",
		"                        S                       ",
		"                   S          S                 ",
		" P                               S    SSSSS     ",
		"                                 K              ",
		"    S             S             SS              ",
		"                      S                         ",
		"                                                ",
	}
}

func level1() []string {
	return []string{
		"                                                                              ",
		"   WS                                                           S             ",
		" SSSS                   SS            S                         K     S       ",
		"                         S                        S      S         S          ",
		"                         S                S                        S          ",
		"D                       SS           S                   S       SS SS        ",
		"                                                SS                 S          ",
		"        S                                 S     SS         S       S          ",
		"                                                                              ",
		"  S                 SSSS S S         SSS                       S              ",
		"             SS    S                                 S         SW         S   ",
		"                   S SSS  SS                S                                 ",
		"                   SPS          S                    S          S       S     ",
		"                                                                              ",
	}
}

func level2() []string {
	return []string{
		"                                                                                                                                  ",
		"                                                                                                                                  ",
		"                                        W                                                             W       W                   ",
		"                                                                                                                                  ",
		"                                                               S                                                                  ",
		"                                                               S                                                                P ",
		"                                                               S                                                                  ",
		" SSSSS  SSSSSSS  SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS SSSSSSSSSS SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS ",
		"           W                                        S            F                                                                ",
		"                                                    S                                                         W                   ",
		"                                                    S                                                               SSS           ",
		"                                        W                                                                                        D",
		"                                                                                                                                  ",
		" SKSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS                                                                          ",
		"                                                                                                                                  ",
		"                                                                                                                                  ",
		"          W                                                      W                                                                ",
		"                                                                                                                                  ",
		"                                                                                                                                  ",
		"                                                                                                                                  ",
	}
}

func level3() []string {
	return []string{
		"                                                                                                                                  ",
		"                                                                                                                                  ",
		"                                                                                                                                  ",
		"          W                                                                                                          F            ",
		"                                                                                                                                  ",
		"                                      S                                                                                           ",
		"      S                                            W        S            K                S                  S                    ",
		"      S        F               W                            S                                                                     ",
		"      S                                                     S    F                                                                ",
		"   SSS SSS         W                                        S                                                                     ",
		"          S     F                                      SSSSSHSSSSS                 S                                              ",
		" P        S                             S                   S                                                        W            ",
		"          S      F                                          S                          W                                          ",
		"   SSS SSS                                        S         S    S                                                                ",
		"      S                     S                               S                                                S                    ",
		"      S    F    F                                                                                                                 ",
		"      S                                                                                                                           ",
		"          W                                                            S                                                          ",
		"                                                                                   F                                              ",
		"                                                                                                                                  ",
		"                                     W                                                      S                          S          ",
		"                        S                         S                                                                               ",
		"                                                                                                                                  ",
		"                                                                         S                                   S                    ",
		"                         W                          F                                         F                                   ",
		"                                                                                                                    W             ",
		"                                                                                                                                  ",
		"                                                              D                                                                   ",
	}
}

func level4() []string {
	return []string{
		"                                                                                                ",
		"                                              SPS        S                                      ",
		"                              F               S S S      S                                      ",
		"                                            SSS S S      S         F                            ",
		"                                           S    S S      S                                      ",
		"                                           S SSSS S      S                      W               ",
		"                                           S      S      S   F    F                             ",
		"                                F           SSSSSS       S                                      ",
		"                                                                                 SSS            ",
		"                                                SSSSSSSSSSS                     S   S           ",
		"                                                S S K S S S                     SS SS           ",
		"                                               S  S S S S  S                    S S S           ",
		"                                               S  S S S S  S                                    ",
		"                                               S  S S S S  S                                    ",
		"                                                                                                ",
		"                                                                                                ",
		"                                                                                                ",
		"              W                                                                                 ",
		"                  SSSSSS SSSSSSSSSSSSSSSSSSSSSS                                                 ",
		"                        S                                                                       ",
		"                        S                                                                       ",
		"                        S              F                                                        ",
		"                        S     W                                     W                           ",
		"                        S                                                                       ",
		"                        S                                                                       ",
		"                        S                                                                       ",
		"              W         S     W                                                                 ",
		"                         S             F                                                        ",
		"                        S S                          H                                          ",
		"                        S  S                                                                    ",
		"                            S                                                                   ",
		"                             S                                                                 D",
		"                              S                                                                 ",
		"                               S                                                                ",
		"                                S                                                               ",
		"                                 F                                                              ",
		"                                  S                                                             ",
		"                                   S               SS                                           ",
		"    F                                             S  S                                          ",
		"               S                                 SSSSSS                                         ",
		"               S           S                       SS                                           ",
		"               SW          S                       SS               F                           ",
		"                           SW                                                                   ",
		"                                                                                                ",
		"                                                                                                ",
		"                                            S                                                   ",
		"                                            S                                                   ",
		"                                            SW                                                  ",
		"                                                                                                ",
		"                                                                                                ",
		"                S                                           S                                   ",
		"                S                                           S                                   ",
		"                SW                                          SW                                  ",
		"                                                                    SSSSSSSSSSSSSSSSSSSSSS      ",
		"                                                                   S                            ",
		"                                                                   S            S        S      ",
		"                           S                                       S              S      S      ",
		"                           S                                       S  H           SW     S      ",
		"    F                      SW                                      S        H            S      ",
		"                                                                    SSSSSSSSSSSSSSSSSSSSS       ",
		"                                                    F                                           ",
		"                                                                                                ",
		"                                                                                                ",
	}
}

func level5() []string {
	return []string{
		"                                                                                                                                                               ",
		"                                                              F                                                                                                ",
		"               SSSS                                                                               SSS                   SSS                 SSSS               ",
		"                   S                                SSSS  SSSS                   F               S   S                 S   S      F             S              ",
		"               S S S                                    SS                                       S S S F               S K                  SHS S              ",
		"               S   S            W                   S S SS S S                                   S   S                 S   S                S   S         W    ",
		"                SSS                                 S   SS   S                                    S S                   SSS                  SS S              ",
		"                                                     SSS  SSS            SSSS                                                                                  ",
		"                                                                             S                                                                                 ",
		"                                                              W          S S S                                SSSS                    SSSS                     ",
		" P                                    SSSS                               S   S                                    S                     H S                   D",
		"                       W                  S                               SSS      F                      W   S S S                     S S                    ",
		"                                      S S S         W                                                         S   S                   S   S                    ",
		"                                      S   S                                                                    S S                     SSS      W              ",
		"                                       SSS                                  W                                                                                  ",
		"                                                                                                                                                               ",
	}
}

func level6() []string {
	return []string{
		"                                                                          D          ",
		" F  HS                    S     S          S                                         ",
		"     K   SSS  SSSS        S               S S                             S          ",
		"   F  S           SS      S     SSSSSS    S                                          ",
		"     S SSSSSSSSS   S      S     S         S                                          ",
		"    S             S      S     S         S                                          ",
		" WW SS     SSSSSSSSSS     S     SF        S                                          ",
		" WWS                S     S     S         S                  WWWWWWWWW               ",
		" WWS                S           S         S                  WWWWWWWWW               ",
		"   S                S           S         S           W      WWWWWWWWW               ",
		"   S                S           SSSSSS    S                                          ",
		"   S                S                     S                                          ",
		"    SSSSSSSSSSSS    S           S  SSSSSSS                                           ",
		" P                  S    SS     S  S                                                 ",
		"                    S    SS     S  SSSSSSSSSSSS                                      ",
		"   SSSSSSSSSSSSSSSSSS           S                                                    ",
		"                                                                                     ",
	}
}

func level7() []string {
	return []string{
		"                                                                                                                                                                                        ",
		"                                                 H         S                 S                                    S     S    K S      S                                                 ",
		"                     S                                                                                 S  H             S    S S          S            S                              HD",
		"    P                                 S          S                           S                                                                                                          ",
		"                                                                                                                                                                                        ",
	}
}

func level8() []string {
	return []string{
		"                    D                           ",
		"                          H                     ",
		"    F                           S   W           ",
		"             W                           S      ",
		"          S        S      S                     ",
		"          S                        S            ",
		"          S                                     ",
		"                        S                       ",
		"                   S          S          H      ",
		" P                  H            S    SSSKS     ",
		"                                 S              ",
		"    S             S             SS              ",
		"                      S                         ",
		"                                                ",
	}
}

func level9() []string {
	return []string{
		"                                                                                              D                                                                                         ",
		"          S                                                                                                                                                                             ",
		"  S       S                                                                                                                                                                             ",
		"   S      S  W                                W                                   SSSS    S   S   SSSS                                                                                  ",
		"    S     S                                                                       S   S    S S    S                                                                                     ",
		"     S    S                                                                       SSSS      S     SSS            S                                              FS S S S S SHS S S S    ",
		"      S   S                                                                       S   S     S     S             S                                                                       ",
		"          S                                                                       SSSS      S     SSSS                                                           S S S S S S S S S S    ",
		"          S                                                                                                                                                                             ",
		"          S              H                                                          H                                       H                                   HS S S S S S K S S S    ",
		"          S                                   W                         S                                                                                                               ",
		"          S S                                                          S                                                                     SSSSSSSSSS          S S S S S S S S S S    ",
		"             S                                                        S                    W                                                              W          F             W    ",
		"              S                                                      S                                                                                           S S S S S S S S S S    ",
		"               S         S                                                                                                                                                              ",
		"            S                                                                                                                                            H       S S S S S SWS S S S    ",
		"                                                                                                                   SSSSSSSSSS                                                           ",
		"                         H                 S              H                    W           W                                                                     S S S S S S S S S S    ",
		"                                            S                                                                                                                                           ",
		"                       S                     S                                                                                                                                  H       ",
		"           W          S                       S                                                                    H                     H                  SSSSSSSSSSSSSSSSSSSSSSSSSSS ",
		"                     S                                                                                                                                     S                    H       ",
		"                    S                                  W                                                                                                   S                            ",
		"                                                                                                                                S                     S    S                            ",
		"                                    SSSSSSSSS                                                  W                                S                      S   S                    S       ",
		"                                         S                                                                                      S                       S  S                            ",
		"    H      W                             S                                                                                      S                        S S        W                   ",
		"                             SSSSSSSSSSSSSSSS                                     S                                             S                          S                            ",
		"                             S                                                     S             SSSSSSSSSSS SSSSSSSSS SSSSSSSSS                           S     S                      ",
		"                             S                                                      S                                           S                          S                            ",
		"                   S      SSSS                         W                W            S                                          S                          S                            ",
		"                  S       S                                                           S                                         S                          S                            ",
		"                 S                             S                                       S                                        S         F                S                            ",
		"                S                               S                                  SSSSSSSSSS                                   S         H           SSSSS                             ",
		"                                                 S                                       S                                      S                          S       SSSSSSSSSSSSSS       ",
		"                                                  S                                       S                                     S                          S       S            S       ",
		"                                                                                           S        F               F           S                          S       S            S       ",
		"                                                                                            S                                   S                          S       S            S       ",
		"                                                                                                                                                           S       SSSSSSSSSSSSSS       ",
		"    H          S            W                                                                                                                              S                            ",
		"              S             W                              SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS                             ",
		"             S              W                     SSSSSSSSS                                                                                                                             ",
		"                            W                    S                        W                                                                                                             ",
		"                                                                                                                                                                                        ",
		"                                      F             SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS                                                             ",
		"                                                   S                                                                                                                                    ",
		"                      S                            S                                                                       SSSSSSSSSS                                                   ",
		"                                                   S                                                                                                                                    ",
		"                                                   S                                                                                                                                    ",
		"             S                                     S                                     W                                                                                              ",
		" SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS                                                                                                                                     ",
		"                                                         S                                                                                                                              ",
		"                                                        S                                                                                                                               ",
		"                                                       S                                                                                          F            S        S               ",
		"             S                        F               S             W                  W                                                          S            S         S              ",
		"                                                     S                   S                                                   S                    S            S                        ",
		"                                                    S                    SS                                                  S                    S            S                        ",
		"                                                   S                      FSSSSSSSSSSSS                                     SS                    S            S                        ",
		"                                                  S                      S                                 W                S                     S                                     ",
		"                                                                         S                                                  S                     S                                     ",
		"                                                       W                 S                                                  S                             S                  W          ",
		"                                                                                                                            S                              S                            ",
		"                      S            S                                W                                                                                       S                           ",
		"                                                                                                                                                                                        ",
		"                                                                                                                                                                                        ",
		"                                                                                                                                                                                        ",
		"                                   S                                                               W                                                           W             W          ",
		"                                                                                                       SS                           S                                                   ",
		"                    S                                                                                 SS                            S                                                   ",
		"                                                                          S                           S        S                     S                                                  ",
		"                                                                         S                                     S                     S                                                  ",
		"  SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS SSSSSSSSSSSSSSSSS      SSSSSSSSSSSSSS SSSSSSSSSSSSSSSSSSSSS SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS  ",
		"                                                                                         HS      SH                                                                                     ",
		"                                                                                       H  S   P  S  H                                                                                   ",
		"                                                                                                                                                                                        ",
	}
}

// demeTombstone is a demo level showing tombstones
func demoTombstoneLevel() []string {
	return []string{
		"                      D                     ",
		"                                            ",
		"                                            ",
		"                                            ",
		"              S                  K          ",
		"                                            ",
		"                                            ",
		"                                            ",
		"                      P                     ",
		"                                            ",
	}
}

// demoWanderLevel is a demo level showing a wander ghost
func demoWanderLevel() []string {
	return []string{
		"                      D                     ",
		"                                            ",
		"                          K                 ",
		"                                            ",
		"              W                             ",
		"                                            ",
		"                                            ",
		"                                            ",
		"                      P                     ",
		"                                            ",
	}
}

// demoFollowLevel is a demo level showing a follow ghost
func demoFollowLevel() []string {
	return []string{
		"                      D                     ",
		"                                            ",
		"                          S          K  S   ",
		"                          S             S   ",
		"              F           SSSSSSSSSSSSSSS   ",
		"                                            ",
		"                                            ",
		"                                            ",
		"                      P                     ",
		"                                            ",
	}
}

func demoHuntLevel() []string {
	// demoHuntLevel is a demo level showing a hunt ghost
	return []string{
		"                      D                     ",
		"                                            ",
		"                          S          K  S   ",
		"                          S             S   ",
		"              H           SSSSSSSSSSSSSSS   ",
		"                                            ",
		"                                            ",
		"                                            ",
		"                      P                     ",
		"                                            ",
	}
}
