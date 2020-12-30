"use strict";

const text = `Input: candidates = [[10,1,2,7,6,1,5]], target = 8, str = "123"
Output: 
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]`;

const itemType = "";
function ioTransfrom() {
  const parts = text.split("Output: ");
  const inputStr = parts[0].replace("Input: ", "");
  const regex = /([a-zA-Z]+ = ([^ ]+))/g;
  const pairs = {};
  let regResult;
  while ((regResult = regex.exec(inputStr))) {
    const pair = regResult[0].endsWith(",")
      ? regResult[0].slice(0, regResult[0].length - 1)
      : regResult[0];
    const [key, value] = pair.split(" = ");
    pairs[key] = eval(value);
  }
  for (const key of Object.keys(pairs)) {
    const value = pairs[key];
    const type = getType(value);
    pairs[key] = parseValue(value, type);
  }
  return Object.keys(pairs)
    .map((key) => `${key}: ${pairs[key]}`)
    .join("\n");
}

function getType(value, path = []) {
  const type = typeof value;
  if (["number", "string"].includes(type)) {
    path.push(type);
    return [...path];
  }
  path.push("array");
  const sample = value[0];
  return getType(sample, path);
}

function getPrefix(path = []) {
  let prefix = "";
  if (path.length > 1) {
    prefix = path
      .map((item) => {
        if (item === "array") {
          return "[]";
        }
        if (itemType) {
          return itemType;
        }
        if (item === "number") {
          return "int";
        }
        if (item === "string") {
          return "string";
        }
      })
      .join("");
  }
  return prefix;
}

function parseValue(value, path) {
  const prefix = getPrefix(path);
  const type = path.shift();
  if (["number", "string"].includes(type)) {
    return value;
  }
  const content = value.map((item) => parseValue(item, [...path])).join(",");
  return `${prefix}{${content}}`;
}
console.log(ioTransfrom());
