function showErrors(errors) {
    var htmlString = ''
    for (var i in errors) {
        htmlString += errors[i]
        htmlString += '<br>'
    }

    var html = HtmlService.createHtmlOutput(htmlString)
        .setTitle('Validation Error')
        .setWidth(300)

    SpreadsheetApp.getUi().showSidebar(html)
}

function isNull(value) {
    if(value == null || value == '') {
        return true
    }
    return false
}

function isNotNull(value) {
    return !(isNull(value))
}

function isItemTextArray(str) {
    var itemText = str.split(/\r\n|\n/)

    for (var i in itemText) {
        if (isNull(itemText[i])) {
            return false
        }

        if (!isItemText(itemText[i])) {
            return false
        }
    }

    return true
}

function isItemText(str) {
    var itemText = str.split(':')

    // type:id:amount の様に、id がある item text 場合
    var len = 3
    if (isNoneId(itemText[0])) {
        len = 2 // type:amount の様に、id がない item text 場合
    }

    if (!itemText.length == len) {
        return false
    }

    return true
}

function isNoneId(type) {
    var noneIdTypes = [
        'no_id_type_1',
        'no_id_type_2',
        'no_id_type_3',
        'no_id_type_4'
    ]

    for (var i in noneIdTypes) {
        if (type == noneIdTypes[i]) {
            return true
        }
    }
    return false
}
  
// format が YYYY/MM/DD かを判定する
function isDate(str) {
    var delim = '/'
    var arr = str.split(delim)
    //Logger.log(arr)
    if (arr.length !== 3) return false

    const date = new Date(arr[0], arr[1] - 1, arr[2])

    if (arr[0] !== String(date.getFullYear()) || arr[1] !== ('0' + (date.getMonth() + 1)).slice(-2) || arr[2] !== ('0' + date.getDate()).slice(-2)) {
        return false
    }

    return true
}

// format が HH:MM:SS かを判定する
function isTime(str) {
    var delim = ':'
    var arr = str.split(delim)
    //Logger.log(arr)
    if (arr.length !== 3) return false

    const date = new Date(2018, 0, 1, arr[0], arr[1], arr[2])

    if (arr[0] !== ('0' + date.getHours()).slice(-2)
        || arr[1] !== ('0' + date.getMinutes()).slice(-2)
        || arr[2] !== ('0' + date.getSeconds()).slice(-2)
    ) {
        return false
    }

    return true
}

// start < end を確認
function isTerm(start, end) {
    var startDate = new Date(start)
    var endDate = new Date(end)

    if (startDate < endDate) {
        return true
    }

    return false
}

// format が YYYY/MM/DD 00:00:00 かを判定する
function isDateTime(str) {
    //Logger.log('isDateTime str = '+str)
    var delim = ' '
    var arr = str.split(delim)
    //Logger.log(arr)
    if (arr.length !== 2) return false

    var resultDate = isDate(arr[0])
    var resultTime = isTime(arr[1])
    //Logger.log('isDate = '+ resultDate +'  isTime = '+ resultTime)

    if (resultDate && resultTime) {
        return true
    }

    return false
}
