### Function syntax
1.

```ts
const myFunction = (param1, param2) => {
  ...
}
```

is equivalent to 

```ts
function myFunction(param1, param2) {
  ...
}
```

2.
Another common function pattern in the code is like this

```ts
const myFunc = (x:number) => (x + 1);
```

is equivalent to 

```ts
const myFunc = (x:number) => {
  return x + 1;
}
```

3. 
Optional parameter:

```ts
const myFunc = (x: number, y:number, z?:number) => {
  //...
}
```

### Destructuring

```ts
const object:TMyType = {
  key1: 10,
  key2: 100,
  key3: 'something else'
};

const myFunc1 = ({key1, key2}:TMyType) => {
  // do something with key1, key2
};

const myFunc2 = (key1:number, key2:number) => {
  // do something with key1, key2
};

myFunc1(object);
myFunc2(object.key1, object.key2);
```

notice the curly bracket in line 7 wrapping `key1, key2`. This is directly extracting the properties `key1, key2` from the parameter. both patterns for `myFunc1` (with destructuring) and `myFunc2` (without destructring) can be used. And it can be a mix of it. e.g.

```ts
const myFunc3(param1:number, param2: string, {key1, key2}:TMyType) {
  // something...
}
myFunc3(10, 'hello', object);
```

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
