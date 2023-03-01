(ns allergies)

(def items-map
  {:eggs 1
   :peanuts 2
   :shellfish 4
   :strawberries 8
   :tomatoes 16
   :chocolate 32
   :pollen 64
   :cats 128})

(defn allergic-to? [score allergy]
  (not= 0 (bit-and score (items-map allergy))))

(defn allergies [score]
  (filter (partial allergic-to? score)
          (keys items-map)))
