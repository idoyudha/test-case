// Case 1
/* Sum of array */
const sumArray = (array) => {
    // console.log("sumArray", array)
    let result = 0
    for (let i = 0; i < array.length; i++) {
        result += array[i]
    }
    return result
}

console.log(sumArray([1, 2, 3]))


// Case 2
/* Date conversion. */
const dateFormatConversion = (string) => {
    let word = string.split(" ")
    let monthNumber = 1
    let monthsName = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
    for (let i = 0; i < monthsName.length; i++) {
        if (monthsName[i] === word[2]) {
            monthNumber += i
        }
    }
    return `${word[3]}-${monthNumber}-${word[1]}`
}

console.log(dateFormatConversion("Tue 16 Jul 2019"))


// Case 3
/* Factorial function */
const calcualteFactorial = (n) => {
    // Base condition from recursive
    if (n === 1) {
        return 1
    }
    return n * calcualteFactorial(n-1)
}

console.log(calcualteFactorial(5))


// Case 4
/* Convert time in numerals into words */
const timeConvert = (h, m) => {

    // Matching number to word function
    const findWord = (h) => {
        timeString = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", 
        "eleven", "twelve", "thirteen", "fourteen", "quarter", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", 
        "twenty-one", "twenty-two", "twenty-three", "twenty-four", "twenty-five", "twenty-six", "twenty-seven", "twenty-eight", "twenty-nine", "thirty"]
        let word = null
        for (let i = 0; i < timeString.length; i++) {
            if (h === i) {
                word = timeString[i - 1]
            }
        }
        return word
    }

    let hour = findWord(h)
    let minutes = findWord(m)
    if (m <= 30) {
        if (m === 0) {
            return `${hour} o'clock`
        } else if (30 > m > 0 && m !== 15) {
            return `${minutes} minutes past ${hour}`
        } else if (m === 15) {
            return `${minutes} past ${hour}`
        } else if (m === 30) {
            return `half past ${hour}`
        }
    } else if (m > 30) {
        let convertMinutes = 60 - m 
        let hourConverted = findWord(h+1)
        let minutesConverted = findWord(convertMinutes)
        if (convertMinutes === 15) {
            return `${minutesConverted} to ${hourConverted}`
        } else {
            return `${minutesConverted} minutes to ${hourConverted}`
        }
    } 
}

console.log(timeConvert(5,28))

// Case 5 
/* Find minimum deletion in array */
const minimumDeletion = (array) => {

    // Create data of number from array with object
    let numbersCounter = {}
    for (let i = 0; i < array.length; i++) {
        if (array[i] in numbersCounter) {
            numbersCounter[array[i]] += 1
        } else {
            numbersCounter[array[i]] = 1
        }
    }

    // Find max counter from data
    let maxCount = 0
    for (let property in numbersCounter) {
        if (numbersCounter[property] > maxCount) {
            maxCount = numbersCounter[property]
        }
    }

    return array.length - maxCount
}

console.log(minimumDeletion([3,3,2,1,3]))

