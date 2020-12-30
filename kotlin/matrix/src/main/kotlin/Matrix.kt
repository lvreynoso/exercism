class Matrix(private val matrixAsString: String) {

    val listRepresentation: List<List<Int>> = destructureStringRepresentation(matrixAsString)

    fun column(colNr: Int): List<Int> {
        val column = listRepresentation.map { row -> row[colNr - 1] }
        return column
    }

    fun row(rowNr: Int): List<Int> {
        return listRepresentation[rowNr - 1]
    }

    private fun destructureStringRepresentation(stringRep: String): List<List<Int>> {
        val rows = stringRep.split("\n")
        val separatedRows = rows.map { rowString -> rowString.split(" ").map { it.toIntOrNull() } }
        val nonNullRows = separatedRows.map { intOrNulls -> intOrNulls.filterNotNull() }
        return nonNullRows
    }
}
