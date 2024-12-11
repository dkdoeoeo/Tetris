package main

const (
	failRotate = 0
	origin     = 1
	up         = 2
	left       = 3
	right      = 4
	left2      = 5
	right2     = 6
	down       = 7
)

const (
	empty       = 0
	typeI       = 1
	typeO       = 2
	typeT       = 3
	typeS       = 4
	typeZ       = 5
	typeL       = 6
	typeJ       = 7
	typeGarbage = 8
	typePreview = 9
)

const (
	Row        = 20
	Col        = 10
	MaxRoomNum = 10
)

type Pos struct {
	x int
	y int
}

// 定義方塊類型，包含中心座標和偏移量
type TetrisBlock struct {
	pos       Pos       //方塊中心座標
	Offsets   [4][2]int // 方塊範圍的相對偏移座標 (共4個格子)
	boardType int
}

var (
	IBlock = TetrisBlock{
		pos: Pos{x: 5, y: 0},
		Offsets: [4][2]int{
			{-1, 0}, // 左邊一格
			{0, 0},  // 中心
			{1, 0},  // 右邊一格
			{2, 0},  // 最右邊
		},
		boardType: 1,
	}

	LBlock = TetrisBlock{
		pos: Pos{x: 6, y: 1},
		Offsets: [4][2]int{
			{-1, 0}, // 左邊一格
			{0, 0},  // 中心
			{1, 0},  // 右邊一格
			{1, -1}, // 右上
		},
		boardType: 6,
	}

	JBlock = TetrisBlock{
		pos: Pos{x: 4, y: 1},
		Offsets: [4][2]int{
			{-1, -1}, // 左上方
			{-1, 0},  // 左邊一格
			{0, 0},   // 中心
			{1, 0},   // 右邊一格
		},
		boardType: 7,
	}

	OBlock = TetrisBlock{
		pos: Pos{x: 5, y: 1},
		Offsets: [4][2]int{
			{0, 0},  // 中心
			{1, 0},  // 右邊一格
			{0, -1}, // 上方一格
			{1, -1}, // 右上角
		},
		boardType: 2,
	}

	ZBlock = TetrisBlock{
		pos: Pos{x: 5, y: 1},
		Offsets: [4][2]int{
			{0, 0},   // 中心
			{1, 0},   // 右邊一格
			{0, -1},  // 上方一格
			{-1, -1}, // 左上角
		},
		boardType: 5,
	}

	SBlock = TetrisBlock{
		pos: Pos{x: 5, y: 1},
		Offsets: [4][2]int{
			{0, 0},  // 中心
			{-1, 0}, // 左邊一格
			{0, -1}, // 上方一格
			{1, -1}, // 右上角
		},
		boardType: 4,
	}

	TBlock = TetrisBlock{
		pos: Pos{x: 5, y: 0},
		Offsets: [4][2]int{
			{0, 0},  // 中心
			{-1, 0}, // 左邊一格
			{1, 0},  // 右邊一格
			{0, 1},  // 下方一格
		},
		boardType: 3,
	}
)
