package main

type Pos struct {
	x int
	y int
}

// 定義方塊類型，包含中心座標和偏移量
type TetrisBlock struct {
	pos     Pos       //方塊中心座標
	Offsets [4][2]int // 方塊範圍的相對偏移座標 (共4個格子)
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
	}

	LBlock = TetrisBlock{
		pos: Pos{x: 6, y: 1},
		Offsets: [4][2]int{
			{0, -1}, // 上方一格
			{0, 0},  // 中心
			{-1, 0}, // 左邊一格
			{-2, 0}, // 最左邊
		},
	}

	JBlock = TetrisBlock{
		pos: Pos{x: 4, y: 1},
		Offsets: [4][2]int{
			{0, -1}, // 上方一格
			{0, 0},  // 中心
			{1, 0},  // 右邊一格
			{2, 0},  // 最右邊
		},
	}

	OBlock = TetrisBlock{
		pos: Pos{x: 5, y: 1},
		Offsets: [4][2]int{
			{0, 0},  // 中心
			{1, 0},  // 右邊一格
			{0, -1}, // 上方一格
			{1, -1}, // 右上角
		},
	}

	ZBlock = TetrisBlock{
		pos: Pos{x: 5, y: 1},
		Offsets: [4][2]int{
			{0, 0},   // 中心
			{1, 0},   // 右邊一格
			{0, -1},  // 上方一格
			{-1, -1}, // 左上角
		},
	}

	SBlock = TetrisBlock{
		pos: Pos{x: 5, y: 1},
		Offsets: [4][2]int{
			{0, 0},  // 中心
			{-1, 0}, // 左邊一格
			{0, -1}, // 上方一格
			{1, -1}, // 右上角
		},
	}

	TBlock = TetrisBlock{
		pos: Pos{x: 5, y: 0},
		Offsets: [4][2]int{
			{0, 0},  // 中心
			{-1, 0}, // 左邊一格
			{1, 0},  // 右邊一格
			{0, 1},  // 下方一格
		},
	}
)
