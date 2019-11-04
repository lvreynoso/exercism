(ns reverse-string)

(defn reverse-string [s] ;; <- arglist goes here
  (apply str (map last (take (count s) (iterate drop-last s))))
)
