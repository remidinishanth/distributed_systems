Ref: https://react-typescript-cheatsheet.netlify.app/docs/basic/getting-started/function_components/ and https://fettblog.eu/typescript-react/components/

### Typescript

#### Mental model:
* Domain of a type = A set of all the values assignable to a variable of that type.
e.g. for boolean type, domain = {true, false}
* Thinking in terms of domains of types helps in reasoning about type operations when we get to the advanced types.

#### TypeScript compiler (`tsc`) performs two independent tasks:
1. Type checking - checks if Domain(type of RHS of the expression) <= Domain(type of LHS of the expression)
2. Compile the TS code to JS code.

Since these tasks are independent, the TS code gets compiled to JS code even if there are type errors. This is a deviation
from other statically typed languages.

### Basics

* `interface` is for when you want to enforce `structural contracts` (i.e what you want passed in or what you want returned back):

### Props

```ts
interface FullName {
    firstName: string;
    lastName: string;
}

function FunctionalComponent(props:FullName){
    // props.firstName
    // props.lastName
}
```

For those cases, you can leverage a JavaScript syntax feature known as destructuring. This allows more flexibility in how you can define your props.

```ts
// Using the same FullName interface from the last example
function FunctionalComponent({firstName, lastName}:FullName){
    // firstName
   // lastName
}
```

What if you wanted to add a middle name? Not everyone has a middle name, so you want to make it optional. Destructuring can provide a clean solution to that problem.

```ts
interface OptionalMiddleName {
    firstName: string;
    middleName?: string;
    lastName: string;
}
function Component({firstName, middleName = "N/A", lastName}:OptionalMiddleName){
    // If middleName wasn't passed in, value will be "N/A"
}
```

### Using React.FC

Another way to define props is to import and use React's `Functional Component` type, FC for short.

Using `React.FC` is more verbose, but does have some added benefits:.

* Explicit with its return type
* Provides type checking and autocomplete for static properties (i.e displayName, defaultProps)
* Provides an implicit definition of children:

```ts
const ReactFCComponent: React.FC<{title:string}> = ({children, title}) => {
    return <div title={title}>{children}</div>
}
```

Ref: https://www.pluralsight.com/guides/defining-props-in-react-function-component-with-typescript and https://stackoverflow.com/questions/59988667/typescript-react-fcprops-confusion


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


### Internationalize React apps

#### ICU Message syntax
Simple message
```
Hello, {name}
```

Complex message
```
Hello, {name}, you have {itemCount, plural,
    =0 {no items}
    one {# item}
    other {# items}
}
```

From a typical JSON file `en.json` containing translation messages:

```json
{
  "helloWorld": "Hello world!",
  "goodbye": "Goodbye"
}
```

```ts
import enMessages from "en.json";

type IntlMessageKeys = keyof typeof enMessages;
```

With react-intl, you can translate messages with the FormattedMessage component or with the formatMessage function. Either way, you have to provide an `id` for the message, which is ordinarily of type `string | number`

```html
<h2>
    <FormattedMessage id="helloWorld"/>
</h2>
<h2>
    {formatMessage({id: "helloWorld"}}
</h2>
```

Ref: https://formatjs.io/docs/react-intl/components#formattedmessage and https://medium.com/weekly-webtips/react-intl-translations-done-properly-with-typescript-3d901ca1b77f

### useState

```ts
// import useState next to FunctionComponent
import React, { FunctionComponent, useState } from 'react';

// our components props accept a number for the initial value
const Counter:FunctionComponent<{ initial?: number }> = ({ initial = 0 }) => {
  // since we pass a number here, clicks is going to be a number.
  // setClicks is a function that accepts either a number or a function returning
  // a number
  const [clicks, setClicks] = useState(initial);
  return <>
    <p>Clicks: {clicks}</p>
    <button onClick={() => setClicks(clicks+1)}>+</button>
    <button onClick={() => setClicks(clicks-1)}>-</button>
  </>
}
```

Ref: https://fettblog.eu/typescript-react/hooks/

Todo: https://javascript.plainenglish.io/how-to-avoid-prop-drilling-in-react-using-component-composition-c42adfcdde1b


### Redux State management

#### Hooks
React Redux provides a pair of custom React hooks that allow your React components to interact with the Redux store.

`useSelector` reads a value from the store state and subscribes to updates, while `useDispatch` returns the store's dispatch method to let you dispatch actions.



Ref: https://react-redux.js.org/introduction/getting-started
