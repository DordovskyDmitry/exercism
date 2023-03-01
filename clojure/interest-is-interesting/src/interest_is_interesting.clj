(ns interest-is-interesting)

(defn interest-rate
  [balance]
  (cond (< balance 0) -3.213
        (< balance 1000) 0.5
        (< balance 5000) 1.621
        :else 2.475) 
  )

(defn interest-rate-in-percent
  [balance]
  (/ (bigdec (interest-rate balance)) 100))

(defn annual-multiplier
  [balance]
  (let [percent (interest-rate-in-percent balance)]
       (cond (> 0 percent) (- 1 percent) 
             :else (+ 1 percent))))

(defn annual-balance-update
  [balance]
  (* balance (annual-multiplier balance)))

(defn amount-to-donate
  [balance tax-free-percentage]
  (cond (>= 0 balance) 0
        :else (int (* balance (/ (* 2 tax-free-percentage) 100)))))