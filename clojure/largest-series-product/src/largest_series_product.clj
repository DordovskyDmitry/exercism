(ns largest-series-product)

(defn- invalid-sequence? [sequence-string]
  (not-every? #(Character/isDigit %) sequence-string))

(defn- to-digits [sequence-string]
  (map #(Character/digit % 10) (sequence sequence-string)))

(defn- largest-valid-product [quantity sequence-string]
  (apply max 
         (map #(apply * %) 
              (partition quantity 1 (to-digits sequence-string)))))

(defn largest-product [quantity sequence-string]
  (cond (invalid-sequence? sequence-string) (throw (Exception. "Invalid string")) 
        (< quantity 0) (throw (Exception. "Invalid quantity"))
        (= quantity 0) 1
        (> quantity 0) (if (empty? sequence-string) 
                         (throw (Exception. "Invalid string"))
                         (largest-valid-product quantity sequence-string))))
