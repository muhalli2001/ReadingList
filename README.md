YO WHATS UP

This will probably be a few months long project but its a side project to help me program a bit better and learn Go in the process.

Right now, all the boilerplate code is done for the API's and I am going through it.
I have routed the API's using HTTP Router. 
They are able to 'marshal' and 'unmarshal' JSON. (basically encode and decode JSON)
The fields of the API requests can be customized, see 'runtime.go' for more information.
And the JSON can get validated. So for example you want a year between 1000 and 3000, you can make that a custom rule with a custom error.
Lots of beneficial stuff so far.

Oh also i didn't describe the project.

You'll be able to view articles from your favorite legacy media.
And you'll be able to view popular books and search for books.

Then you'll be able to add these articles and books to 'playlists' i.e. readinglists, to document what you've been reading
Just a way to make reading a little more accessible and easy. Media makes easy objects to focus on when building API's since there are lots of open source products that provide resources for getting books and articles.

These are for the APIs which should include:

1. Book searches provided by Open Library.
2. Book cover images also provided by Open Library.
3. List of Articles from X legacy media company.
4. List of Articles from Y legacy media company.
5. List of Articles from Z legacy media company.
6. List of Articles from etcetera.
7. CRUD functions for users to make accounts, see their accounts, update their accounts and delete their accounts.

Really that's all i've got so far I might be missing some things.
But actionable next step is to create a feature branch and work on one of these API's!
