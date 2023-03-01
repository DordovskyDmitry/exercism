(ns robot-simulator)

(defn robot [coordinates, direction]
  {:coordinates coordinates :bearing direction})

(def turn-right
  {:north :east
   :east :south
   :south :west
   :west :north})

(def turn-left
  {:north :west
   :west :south
   :south :east
   :east :north})

(defn- update-params [pos]
  (case (pos :bearing)
    :east [[:coordinates :x] inc]
    :west [[:coordinates :x] dec]
    :north [[:coordinates :y] inc]
    :south [[:coordinates :y] dec]))

(defn- apply-step [pos step] 
  (case step
    \R (update pos :bearing turn-right)
    \L (update pos :bearing turn-left)
    \A (apply update-in pos (update-params pos))))

(defn simulate [path, pos]
  (reduce apply-step pos path))
