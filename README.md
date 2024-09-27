This was created in order to have an overview of my spending on freshful.ro.
Not everything is automated(yet)!

Create .env file based on the example in the repo. Go to your browser on freshful.ro and open developer tools and search in network tab 
Now browse to "Comenzile mele". Search for requests matching for any "shop/order" and inspect the header from there to copy the Bearer token. Put in the .env file

The JWT is valid for 30 minutes usually. Now you can run 

    make download

Now you will have a ./data folder containing a json for each order you made in freshful.
In order to do any bare analysis I opted to denormalize all order in csv format (1 line per order line). To do that run:

    make denormalize

Now you will have `denormalized.csv` available. I usually start doing pivot tables over that data.[README.md](README.md)