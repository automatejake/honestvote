package p2p

// if tran, ok := block.Transaction.(primitive.D); ok {
// 	tranMap := tran.Map()

// 	if pos, ok := tranMap["positions"].(primitive.A); ok {
// 		var posElements primitive.A
// 		tranMap["positions"] = nil

// 		for _, position := range pos {
// 			if posInfo, ok := position.(primitive.D); ok {
// 				posMap := posInfo.Map()

// 				if cand, ok := posMap["candidates"].(primitive.A); ok {
// 					var candElements primitive.A
// 					posMap["candidates"] = nil

// 					for _, candidate := range cand {
// 						if candInfo, ok := candidate.(primitive.D); ok {
// 							candMap := candInfo.Map()
// 							candElements = append(candElements, candMap)
// 						}
// 					}

// 					posMap["candidates"] = candElements
// 				}

// 				posElements = append(posElements, posMap)
// 			}
// 		}

// 		tranMap["positions"] = posElements
// 	}

// 	block.Transaction = tranMap
// }
