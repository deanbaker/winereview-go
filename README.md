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

### Notes

There are many ways to skin a cat, as they say. This is just one potentail solution to the problem. You might also have noticed that I have created a new package called **wine**. This is an example of how you can and probably should start packaging your code. 

There is also a bug (feature?) when you start to look at the file with 130k wine reviews. Did you catch it? THe most prolific reviewer has no name, what would you do to improve the code to catch this?

## Step 2

Now we have some code that can ingest a csv file, and collate some statistics about it. How about we look at exposing a RESTful api to play the data back?

    **Note that we will only be using the golang standard libraries, please do not import any external dependencies for this task**

### The Challenge
We will create a full CRUD experience for the reviews we have collated. For this challenge we will not worry about pagenation, that will come in a later task. For now we should be happy to play back all the reviews we have in the system.

You can see the Open API Specification under **/resources/swagger-step2.yaml** - head to https://editor.swagger.io/ and have a look.


I would look at the https://golang.org/pkg/net/http/ package for inspiration.

## Notes

Your swagger editor should now be able to make calls into your local machine. Note how we have had to set appropriate headers on each HandleFunc function to cater for this. Is there a better way? 

Inspiration for the package layout was taken from this [Meduim](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) article, I saw it referenced in this rather entertaining talk on [YouTube](https://www.youtube.com/watch?v=MzTcsI6tn-0). Have a look if you get a chance and see where I was coming from.

## Step 3

Now we have a functioning RESTful API, it would be great to use something other than an in memory datastore to get things going.
Notice that in Step 2 we had abstracted out some interfaces to deal with working with the Review struct?

### The Challenge

Implement these functions to store data into a database. By reusing the existing interfaces this task should require quite a small touch


