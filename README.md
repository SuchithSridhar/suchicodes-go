# Suchicodes In Go + Templ + HTMX

I'm trying to rewrite [suchicodes](https://suchicodes.com) using Go, Templ, and
HTMX.

I decided on this because it was due for a rewrite (as I've grown as a
dev) and after a lot of research this seemed to be the most popular and
effective solution. This isn't with the focus of trying to write something as
fast as possible but rather trying to learn Go while paying as little as
possible for my website hosting.

However, I found that the documentation was good but maybe due to lack of
experience or other reasons, I was stuck on particular issues when trying to
build it from scratch. Some pretty basic things were bothering this approach.

I may retry this using `html/template` since it looks like it would work a
little better than `templ`, in the sense that it would be easier to do things
I'm trying to do.

Here are some problems I'm having:

- **Login/Logout management**: I want to use sessions to manage logged in users.
  However, most online resources on Go with Echo (framework) do not address how
  I can set in each request (using middleware) on whether a user is logged in.
  This is still manageable with a lot of redundant code (in my view).

- **Theme management**: Similar to login and logout, I want a user to be able to
  set a theme for the site which is persistent even when the browser is closed.

- **Templating**: The way templating in Jinja is done is using a sort of
  placeholder for `body` and `head` where stuff and then be inserted. There is a
  way to do this using `children...` in Templ but I'm not sure how I'd do this
  for multiple parts.

The hardest part of this is that I have such a hard time finding an example of a
GOOD website built using these technologies. The closest thing I could find to
this was: <https://github.com/muhwyndhamhp/marknotes> but that was not similar
enough to help me.


## Update 2023, Dec 31 (3:17 am)

Having spent way too much time on this, I finally figured out how to set
cookies and update sessions. I also figured out how I can set contexts and
variables that can be accessed by `templ` directly. There was so research and
testing that had to be done... I'm so tired so I'ma head to bed now.
