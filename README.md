# Wine review API in Go

## Step 1

### Data
We have found a set of data from [Kaggle](https://www.kaggle.com/zynicide/wine-reviews#winemag-data-130k-v2.csv) that we have downloaded into the /resources directory. There are two files, one with 130k records, and another with 20 records. I would suggest working with the csv with 20 records for when developing.

Have a look at the kaggle link above, and check out the csv format - in particular we will be interested in the following attributes:

* ID
* Title
* Variety
* Country
* Points
* TasterName

### The Challenge
Create a go program that will read a csv file and print out the following information:

```
2019/03/31 13:30:35 Starting application
2019/03/31 13:30:35 Number of reviews: 20
2019/03/31 13:30:35 Most prolific Reviewer: Michael Schachner
2019/03/31 13:30:35 Most reviewed variety: Gewürztraminer
2019/03/31 13:30:35 Finished the application
```

Run the program by simply use the following command
``` sh
go run main.go
```

