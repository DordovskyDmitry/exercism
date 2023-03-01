(ns cars-assemble)

(defn success-rate
  [speed]
  (cond (= speed 0) 0.0
        (< speed 5) 1.0 
        (< speed 9) 0.9 
        (= speed 9) 0.8 
        (= speed 10) 0.77))

(def base-productivity 221)

(defn production-rate
  "Returns the assembly line's production rate per hour,
   taking into account its success rate"
  [speed]
  (* base-productivity speed (success-rate speed)))

(defn working-items
  "Calculates how many working cars are produced per minute"
  [speed]
  (int (/ (production-rate speed) 60)))
