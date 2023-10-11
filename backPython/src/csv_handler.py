def getAllCallsOfType(cdr_csv, type):
    return len(cdr_csv[cdr_csv["type"]==type])


def getAnsweredCallsOfType(cdr_csv, type):
    return len(cdr_csv[(cdr_csv["type"]==type) & (cdr_csv["status"]=="ANSWERED")])
    