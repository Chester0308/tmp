Validator = function(sheet, rules) {
    this.sheet = sheet;
    this.rules = rules;
    
    // 1行目は colum 名
    var tmp = sheet.getRange(1, 1, 1, sheet.getLastColumn()).getValues();
    this.columns = tmp[0]

    // 2行目から取得する
    this.rows = sheet.getRange(2, 1, sheet.getLastRow() - 1, sheet.getLastColumn()).getValues();

    this.functions = {
        'not_null': { func: isNotNull, msg: 'is required' },
        'datetime': { func: isDateTime, msg: 'is not datetime' },
        'item_text_array': { func: isItemTextArray, msg: 'is not valid' }
    }
}
  
Validator.prototype.getColumnIndex_ = function(column) {
    for (var i in this.columns) {
        if (this.columns[i] == column) {
            return i
        }
    }

    return 0
}

Validator.prototype.run = function() {
    var errors = []

    for (var column in this.rules) {
        var x = this.getColumnIndex_(column)

        for (var i in this.rules[column]) {
            var type = this.rules[column][i]

            for (var y in this.rows) {
                var val = this.rows[y][x]
                //Logger.log('column:%s  val: %s', column, val)
                
                if (!this.functions[type].func.call({}, val)) {
                    var msg = this.functions[type].msg
                    var rowNumber = parseInt(y) + 2 // index のため +1 と、1行目(column 名)の分 +1 をあわせて +2
                    errors.push(column +' '+ msg +' (row:'+ rowNumber +'  value:'+ val +')')
                }
            }
        }
    }

    if (errors.length > 0) {
        showErrors(errors)
        return false
    }

    return true
}

// 配列の値に重複がないかチェックする
Validator.prototype.validationUnique = function(columns) {
    var errors = {}
    var values = this.rows
    var tmp = {}
    for (var i in columns) {
        tmp[columns[i]] = {}
    }

    for (var i in values) {
        for (var j in columns) {
            var cIndex = this.getColumnIndex_(columns[j])
            var val = values[i][cIndex]
            //Logger.log('columns:%s   cIndex:%s', columns[j], cIndex)
            //Logger.log(val)
            
            if (tmp[columns[j]][val]) {
                var rowNumber = parseInt(i) + 2
                errors.push(columns[j] +' is not unique (row:'+ rowNumber +'  value:'+ val +')')
            }
            tmp[columns[j]][val] = 1
        }
    }

    if (errors.length > 0) {
        showErrors(errors)
        return false
    }

    return true
}
