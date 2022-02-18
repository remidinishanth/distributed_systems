### Children are props

Chances are if you've written React before, you've dealt with props and children in some way. Let's say we have a super simple button component:

```ts
const Button = () => (
  <button>
    I am a button.
  </button>
)
```

If you want to pass things to this button, you would use a prop.

```ts
// our button
const Button = ({ color }) => (
  <button className={color}>
    I am a button
  </button>
)

// somewhere else
<Button color="red" />
```

If you want to make our button say more than just "I am a button," you can pass children to it.

```ts
// our button
const Button = ({ color, children }) => (
  <button className={color}>
    {children}
  </button>
)

// somewhere else
<Button color="red">
  I am still a button
</Button>
```

By passing children in this way, you are passing it to the component by position. Now, if you notice that little header there of this section, I call children a prop. Did you know that it can be passed as a named prop, too?

```ts
// turn this
<Button color="red">
  I am still a button
</Button>

// into this
<Button color="red" children={"I am still a button"} />
```

These two syntaxes produce the exact same result on the page! Children is a prop, and can be passed in to components in different ways.

Ref: https://www.netlify.com/blog/2020/12/17/react-children-the-misunderstood-prop/
