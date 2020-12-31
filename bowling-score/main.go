// https://codezine.jp/article/detail/13320
// ボーリングスコア
// 回答では127だが、これは模範解答が間違っている気がする
// 手動で計算しても137が解答なので、多分9フレームのストライクに対しての加算が計算されていない
package main

import (
	"fmt"
	"log"
)

type Frames [][]int

var game1 = Frames{
	{6, 4}, {8, 0}, {10}, {2, 7}, {5, 5}, {3, 4}, {10}, {9, 1}, {1, 2}, {7, 1},
}
var game2 = Frames{
	{1, 8}, {9, 1}, {7, 2}, {10}, {0, 0}, {9, 1}, {3, 6}, {8, 0}, {5, 4}, {10, 8, 1},
}
var game3 = Frames{
	{10}, {10}, {10}, {10}, {10}, {10}, {10}, {10}, {10}, {10, 10, 10},
}
var game4 = Frames{{9, 1}, {8, 2}, {10}, {5, 0}, {3, 6}, {4, 2}, {7, 3}, {6, 3}, {10}, {9, 1, 9}}

func main() {
	if err := Run(); err != nil {
		log.Fatalln(err)
	}
}

func Run() error {
	for i, points := range []Frames{game1, game2, game3, game4} {
		score, err := CalculateScore(points)
		if err != nil {
			return err
		}
		fmt.Printf("%d: %d\n", i, score)
	}
	return nil
}

func CalculateScore(frames Frames) (int, error) {
	// フレームごとのスコア
	var scores []int

	// フレームごとに計算していく
	for frameIndex, frame := range frames {
		// このフレーム分の得点を単純に足す
		var currentFrameScore int
		for _, score := range frame {
			currentFrameScore += score
		}

		// 最後のフレーム以外は追加加算あり
		if frameIndex < 10 {
			var additionalScore int

			// ストライク: 次の2投分のスコアを加算する
			if frame[0] == 10 {
				nextFrameIndex := frameIndex + 1
				// 次のフレームの1投目、2投目を足す
				if len(frames) > nextFrameIndex {
					additionalScore += frames[nextFrameIndex][0]
					if len(frames[nextFrameIndex]) > 1 {
						additionalScore += frames[nextFrameIndex][1]
					} else if len(frames) > nextFrameIndex+1 {
						additionalScore += frames[nextFrameIndex+1][0]
					}
				}
			} else if frame[0]+frame[1] == 10 {
				// スペア: 次のフレームの1投目のスコアを足し合わせる
				nextFrameIndex := frameIndex + 1
				if len(frames) > nextFrameIndex {
					additionalScore += frames[nextFrameIndex][0]
				}
			}

			currentFrameScore += additionalScore
		}

		scores = append(scores, currentFrameScore)
	}

	// スコア合計
	var totalScore int
	for _, frameScore := range scores {
		totalScore += frameScore
	}
	return totalScore, nil
}
