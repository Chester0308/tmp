function onOpen() {
    // Add a custom menu to the active document, including a separator and a sub-menu.
    var ui = SpreadsheetApp.getUi();
    ui.createMenu('Validation Menu')
        .addItem('First item', 'runValidation')
        .addSeparator()
        .addSubMenu(ui.createMenu('Sub-menu')
        .addItem('Second item', 'menuItem2'))
        .addToUi();
}

function runValidation() {
//    runEventIndexMaster()
    runRaidRankBonusMaster()
}

function runEventIndexMaster() {
    var sheet = SpreadsheetApp.getActiveSheet();
    var rules = {
        'id': ['not_null'],
        'open_date': ['not_null', 'datetime'],
        'close_date': ['not_null', 'datetime'],
        'result_close_date': ['not_null', 'datetime']
    }
    var uniqueColumns = [
        'id'
    ]

    var val = new Validator(sheet, rules);

    var isValied = val.validationUnique(uniqueColumns)
    Logger.log('is Unique = '+ isValied)

    isValied = (isValied && val.run())
    Logger.log('run Validation = '+ isValied)
}

function runRaidRankBonusMaster() {
    var sheet = SpreadsheetApp.getActiveSheet();
    var rules = {
        'id': ['not_null'],
        'present': ['not_null', 'item_text_array']
    }
    var uniqueColumns = [
        'id'
    ]

    var val = new Validator(sheet, rules);

    var isValied = val.validationUnique(uniqueColumns)
    Logger.log('is Unique = '+ isValied)

    isValied = (isValied && val.run())
    Logger.log('run Validation = '+ isValied)
}

function menuItem2() {
    SpreadsheetApp.getUi().alert('You clicked the second menu item!')
}
