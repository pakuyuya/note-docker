UPDATE sample.musers SET disabled=TRUE, updateat=CURRENT_TIMESTAMP, upduser='sample-batch' WHERE expiredt <= CURRENT_DATE AND NOT disabled;
