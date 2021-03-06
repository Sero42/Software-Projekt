package constants

const (
	MaxWinSizeX = 1000
	MaxWinSizeY = 700

	// Charactertypen
	WhiteBomberman = 1 // Spielfiguren im Single Player Mode
	BlackBomberman = 2
	BlueBomberman  = 3
	RedBomberman   = 4
	WhiteBattleman = 5 // Spielfiguren im Multi Player Mode
	BlackBattleman = 6
	BlueBattleman  = 7
	RedBattleman   = 8
	Balloon        = 9
	Teddy          = 10
	Ghost          = 11
	Drop           = 12
	Pinky          = 13
	BluePopEye     = 14
	Jellyfish      = 15	// stufe 2
	Snake          = 16
	Spinner        = 17
	YellowPopEye   = 18
	Snowy          = 19
	YellowBubble   = 20
	PinkPopEye     = 21
	Fire           = 22	// stufe 3
	Crocodile      = 23
	Coin           = 24
	Puddle         = 25
	PinkDevil      = 26
	Penguin        = 27
	PinkCyclops    = 28
	RedCyclops     = 29
	BlueRabbit     = 30
	PinkFlower     = 31
	BlueCyclops    = 32
	Fireball       = 33
	Dragon         = 34
	BlueDevil      = 35

	//Character attributes:
	CBoxSize = 15 // Size of collison Box

	//Items:
	Bomb            = 100
	PowerItem       = 101 // increases bomb power
	BombItem        = 102 // ability to put several bombs
	PunchItem       = 103
	HeartItem       = 104 // remote detonator for bombs
	RollerbladeItem = 105 // increases player's speed
	SkullItem       = 106 // decreases player's lifes
	WallghostItem   = 107 // lets player walk through destroyable tiles
	BombghostItem   = 108 // lets player walk over bombs
	LifeItem        = 109 // increases lifes by 1
	Exit            = 110
	KickItem        = 111 // ability to kick bombs away

	//Wände bzw. Hindernisse zum Wegsprengen:
	Stub           = 120
	Brushwood      = 121
	Greenwall      = 122
	Greywall       = 123
	Brownwall      = 124
	Darkbrownwall  = 125
	Evergreen      = 126
	Tree           = 127
	Palmtree       = 128
	Perl           = 129
	Snowrock       = 130
	Greenrock      = 131
	House          = 132
	Christmastree  = 133
	Perl1          = 134
	Perl2          = 135
	Perl3          = 136
	Perl4          = 137
	Littlesnowrock = 138

	// Bewegung bzw. Zustand eines Characters
	Stay  = 0
	Up    = 1
	Down  = 2
	Left  = 3
	Right = 4
	Dead  = 5
	Intro = 6

	// Musik
	ThroughSpace        = 1
	TheFieldOfDreams    = 2
	OrbitalColossus     = 3
	Fight               = 4 // Super für Boss-Level
	JuhaniJunkalaTitle  = 5
	JuhaniJunkalaLevel1 = 6
	JuhaniJunkalaLevel2 = 7
	JuhaniJunkalaLevel3 = 8
	JuhaniJunkalaEnd    = 9
	ObservingTheStar    = 10
	MyVeryOwnDeadShip   = 11

	//Sounds
	Deathflash = 100
	Falling1   = 101
	Falling2   = 102
	Falling3   = 103
	Falling4   = 104
	Falling5   = 105
	Falling6   = 106
	Falling7   = 107
	Falling8   = 108
	Falling9   = 109
	Falling10  = 110
	Falling11  = 111
	Falling12  = 112
	Alarm1     = 113
	Alarm2     = 114
	Alarm3     = 115
	Alarm4     = 116
	Alarm5     = 117
	Alarm6     = 118
	Alarm7     = 119
	Alarm8     = 120

	// Arena Details
	TileSize   = 16
	WallWidth  = 32
	WallHeight = 16

	// Arenas
	MfS        = 0
	TurfNtrees = 1
	Castle     = 2

	// Eigenschaften einer Position im Spielfeld (Tile)
	Free          = 0
	Undestroyable = 1
	Destroyable   = 2
	Blocked       = 3
)
