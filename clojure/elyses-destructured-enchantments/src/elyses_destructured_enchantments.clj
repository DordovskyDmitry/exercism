(ns elyses-destructured-enchantments)

(defn first-card
  "Returns the first card from deck."
  [deck]
  (let [[item1] deck] item1))

(defn second-card
  "Returns the second card from deck."
  [deck]
  (let [[_ item2] deck] item2))

(defn swap-top-two-cards
  "Returns the deck with first two items reversed."
  [deck]
  (let [[item1 item2 & rest] deck]
    (conj rest item1 item2)))

(defn discard-top-card
  "Returns a sequence containing the first card and
   a sequence of the remaining cards in the deck."
  [deck]
  (let [[item1 & rest] deck] [item1, rest]))

(def face-cards
  ["jack" "queen" "king"])

(defn insert-face-cards
  "Returns the deck with face cards between its head and tail."
  [deck]
  (let [[item1 & rest] deck]
    (cond (nil? item1) face-cards 
          (nil? rest) (apply vector (cons item1 face-cards))
          :else (apply vector (concat [item1] face-cards rest)))))
