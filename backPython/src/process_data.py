import pandas as pd
from io import StringIO
import csv_handler


def process(stringCSV):

    stringCSV = StringIO(stringCSV)
    csv = pd.read_csv(stringCSV, delimiter=',')
    # csv.to_csv("test.csv")
    inbounds = csv_handler.getAllCallsOfType(csv, "Inbound")
    outbounds = csv_handler.getAllCallsOfType(csv, "Outbound")
    inbounds_answered = csv_handler.getAnsweredCallsOfType(csv, "Inbound")
    outbounds_answered = csv_handler.getAnsweredCallsOfType(csv, "Outbound")

    return inbounds, outbounds, inbounds_answered, outbounds_answered

