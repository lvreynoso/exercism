import kotlin.math.sqrt

object Sieve {

    fun primesUpTo(upperBound: Int): List<Int> {
        val nonprimes: Array<Boolean> = Array(upperBound) { true }
        nonprimes[0] = false
        val lastTestInteger = sqrt(upperBound.toDouble()).toInt()
        for (i in 2..lastTestInteger) {
            var multiplier = i
            var product = i * i
            while (product <= upperBound) {
                nonprimes[product - 1] = false
                multiplier += 1
                product = i * multiplier
            }
        }
        val primes = nonprimes.indices.filter { i -> nonprimes[i] }
        return primes.map { int -> int + 1 }
    }
}
