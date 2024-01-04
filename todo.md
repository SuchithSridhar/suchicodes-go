# ToDo List

- [x] Update the theme changer callback to callback to a function to update the
      icons for Github and LinkedIn in the footer.

- [x] Update Navbar to collapse when the device width is too small.

- [x] Update container div to expand as the device width gets smaller.

- [ ] Set an upper bound on the width for the container div for things like wide
      monitors.

- [ ] Set up the database interactions and set up functions to access database
      information.

- [ ] Set up the markdown renderer and set up functions to convert markdown to
      HTML. Decide between creating objects each time vs retaining objects
      permanently.

- [ ] Don't use the TOC extension write your own parser so that it can be used
      in a custom way. Double check the IDs that are given to each of the
      headers.

- [ ] Set up a simple way to generate UUIDs and common operations on UUIDs.

- [ ] Set up a way to hash passwords and compare passwords using bcrypt.


Here is a list of end-points that have to be built:

- [ ] /projects
- [ ] /projects/<project id>
- [ ] /url and /u
- [ ] /blogs (maybe have multiple "tabs" to view particular forms -> HTMX)
- [ ] /blogs/<blog id>


Ideas for /blogs:

[ Tree View | Card View | Search ]

filter: category: [...] search: [...] sort: [...]

The search and sort can only be something that works for card view.
To view all blogs under a category, we head to card view and filter by that
particular category. (this filter by category should be a dropdown but still
input field so that you can either input text or select from the dropdown.) The
filter in the card view should manage filtering by tags, author, date etc?.

The search view can just be a search bar that searchs all possible fields with
checkboxes for what to search on. It can also filter by category and maybe even
by tags?
