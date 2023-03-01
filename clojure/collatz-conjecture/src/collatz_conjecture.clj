(ns collatz-conjecture)

(defn- collatz-with-acc [num, steps]
  (cond (= num 1) steps
        (odd? num) (collatz-with-acc (inc (* 3 num)) (inc steps))
        :else (collatz-with-acc (/ num 2) (inc steps))))

(defn collatz [num] 
  (if (< num 1) 
    (throw (Exception. "Invalid input value"))
    (collatz-with-acc num 0)))
