import kotlin.math.sqrt

object Sieve {

    fun primesUpTo(upperBound: Int): List<Int> {
        val nonprimes: MutableSet<Int> = mutableSetOf<Int>()
        val lastTestInteger = sqrt(upperBound.toDouble()).toInt()
        for (i in 2..lastTestInteger) {
            var multiplier = i
            var product = i * i
            while (product <= upperBound) {
                nonprimes.add(product)
                multiplier += 1
                product = i * multiplier
            }
        }
        val primes = (2..upperBound).filter { !nonprimes.contains(it) }
        return primes
    }
}
