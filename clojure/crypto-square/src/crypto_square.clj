(ns crypto-square
  (:require [clojure.string :as str]))

(defn normalize-plaintext [txt]
  (str/replace (str/lower-case txt) #"[^\d\w]" ""))

(defn square-size [txt] 
  (-> txt count Math/sqrt Math/ceil int))

(defn plaintext-segments [txt]
  (let [normalized (normalize-plaintext txt)]
    (map (partial apply str) 
         (partition-all (square-size normalized) 
                        normalized))))

(defn- equal-len-segments [words]
  (let [len (count (first words))]
    (map #(format (str "%-" len "s") %) words)))

(defn- transposed-segments [txt]
  (->> txt plaintext-segments equal-len-segments (apply (partial map str))))

(defn ciphertext [txt]
  (str/replace (apply str (transposed-segments txt)) " " ""))

(defn normalize-ciphertext [txt] 
  (str/join " " (transposed-segments txt)))
