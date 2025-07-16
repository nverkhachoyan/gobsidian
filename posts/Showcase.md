[X] Fix the issue where a link to a post inside nested folders doesn't resolve to the post if the link only contains the name like so

[[Personal/Tomfoolery And Stuff/What are threads anyway?]]


$$ \frac{a}{b} $$

$$\int_0^1 x^2 dx$$ 

```mermaid
sequenceDiagram
Alice->>John: Hello John, how are you?
loop HealthCheck
    John->>John: Fight against hypochondria
end
Note right of John: Rational thoughts!
John-->>Alice: Great!
John->>Bob: How about you?
Bob-->>John: Jolly good!
 
```


```html
{{if .Tags}}

<div class="left-sidebar-item">

<h5 class="left-sidebar-item-title">Tags</h5>

<div class="left-sidebar-item-content">

<nav class="tags-nav">

<ul>

{{range .Tags}}

<li><a href="/tag/{{.Slug}}/">#{{.Name}}</a></li>

{{end}}

</ul>

</nav>

</div>

</div>

{{end}}
```


>[!info]
>This info callout is great. It is informative as heck!

  

>[!warning]
>This warning callout is warning you about the monster under your bed!

  

>[!success]
>Great success!
>
>**Lorem Ipsum** is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.

