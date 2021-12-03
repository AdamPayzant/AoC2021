import System.IO  
import Control.Monad
import Data.List

main = do  
        let list = []
        handle <- openFile "input" ReadMode
        contents <- hGetContents handle
        let singlewords = words contents
            list = f singlewords
        print (sweepWindows list)
        hClose handle   

f :: [String] -> [Int]
f = map read

-- Part 1
sweep :: [Int] -> Int

sweep [_] = 0
sweep (x:y:xs) = (if x < y then 1 else 0) + sweep (y:xs)

-- Part 2
sweepWindows :: [Int] -> Int

sweepWindows l = sweep (map g (tails l))
    where
    g (x:y:z:_) = x+y+z
    g _ = 0

