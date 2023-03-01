(ns difference-of-squares)

(defn sum-of-squares [n]
  (/ (* n (inc n) (inc (* 2 n))) 6))

(defn square-of-sum [n]
  (/ (* n n (inc n) (inc n)) 4))

(defn difference [n]
  (- (square-of-sum n) (sum-of-squares n)))