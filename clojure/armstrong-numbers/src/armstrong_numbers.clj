(ns armstrong-numbers)

(defn order [x]
  (->> x
       (iterate #(quot % 10))
       (take-while pos?)
       (count))
)

(defn digits [n]
  (->> n
       (iterate #(quot % 10))
       (take-while pos?)
       (mapv #(mod % 10))
       rseq)
)

(defn exp [x n]
  (reduce * (repeat n x))
)

(defn armstrong? [num] ;; <- arglist goes here
  (cond 
    (< num 10) true
    :else (= num (reduce + (mapv #(exp % (order num)) (digits num))))
  )
)
