package helper

func CountEstimatedTimeToFix(typeOfDamage string, hasBeenDamagedBefore bool) int {
    if hasBeenDamagedBefore {
        switch typeOfDamage {
        case "Rusak Sebagian":
            return 4
        case "Rusak Sedang":
            return 8
        case "Rusak Parah":
            return 14
        }
    } else {
        switch typeOfDamage {
        case "Rusak Sebagian":
            return 2
        case "Rusak Sedang":
            return 4
        case "Rusak Parah":
            return 7
        }
    }

    // Return a default value if none of the conditions are met
    // This could be an indication of an error, such as an unrecognized typeOfDamage
    return 0
}
