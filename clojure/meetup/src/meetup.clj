(ns meetup
  (:import [java.time LocalDate]
           [java.time YearMonth]))

(def ord->index
  {:first 0
   :second 1
   :third 2
   :fourth 3})

(def weekday->index
  {:monday 1
   :tuesday 2
   :wednesday 3
   :thursday 4
   :friday 5
   :saturday 6
   :sunday 7})

(defn- days-count-in-month [month year]
  (.lengthOfMonth (YearMonth/of year month)))

(defn- weekday [month year n]
  (-> (LocalDate/of year month n) .getDayOfWeek .getValue))

(defn- weekdays-in-month [month year]
  (map #(weekday month year (inc %))
       (range (days-count-in-month month year))))

(defn- detect-date [ord dates] 
  (cond (= ord :teenth) (first (filter #(<= 13 % 19) dates)) 
        (= ord :last) (last dates) 
        :else (nth dates (ord ord->index))))

(defn- date [month year weekday ord]
  (let [weekday-index (weekday weekday->index)]
    (->> (weekdays-in-month month year)
         (keep-indexed #(when (= %2 weekday-index) (inc %)))
         (detect-date ord))))
  
(defn meetup [month year weekday ord]
  [year month (date month year weekday ord)])
