Tolerates the loss of a minority of ensemble members and still function.
* To tolerate a loss of `n` members, we need atleast `2 * n + 1` nodes because for `n` to be the minority, we need atleast `2 * n + 1` nodes.
* It's recommended to have odd number(3 or 5) of nodes because we want to have majority surviving to continue to function, You don't get any benefit by having 6 nodes instead of 5 nodes, for both 5 or 6 nodes, we can only have loss of 2 nodes.
