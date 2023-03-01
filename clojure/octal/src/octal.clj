(ns octal)

(def base 8)

(defn char->digit [d]
  (Character/digit d 10))

(defn- exp [base power]
  (reduce * (repeat power base)))

(defn- valid? [s]
  (every? #(<= 0 (char->digit %) (dec base)) s))

(defn to-decimal [s] 
  (if (valid? s)
    (reduce + (map-indexed #(* (char->digit %2) 
                               (exp base %)) 
                           (reverse s)))
    0))
  
