(ns bob
  (:require [clojure.string :as str]))

(defn response-for [s]
  (cond
    (str/blank? s) "Fine. Be that way!"
    :else (let [l (-> s 
                       (str/replace #"[A-Z]" "U")
                       (str/replace #"[0-9]" "N")
                       (str/replace #"[a-z]" "L"))]
            (cond
              (or (str/includes? l "L") (and (str/includes? l "N") (not (str/includes? l "U")))) 
                (cond 
                  (str/ends-with? l "?") "Sure."
                  :else "Whatever.")
              :else (cond 
                      (str/ends-with? l "?") "Calm down, I know what I'm doing!"
                      :else "Whoa, chill out!"))))
)
