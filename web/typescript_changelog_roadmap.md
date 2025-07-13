# 7.0 TBA

[First Release with GO Compiler](https://devblogs.microsoft.com/typescript/typescript-native-port/)

https://github.com/microsoft/typescript-go

# 6.0 TBA

> some deprecations and breaking changes to align with the upcoming native codebase

# 5.9 TBA

# 5.8 (Feb 2025)

* **Performance/Editor improvements** — fewer array allocations during path normalization, faster program loading & watch-mode updates ([typescriptlang.org][5]).
* **Stable `--module node18`** — ESM-friendly mode for Node.js 18 that:

  * Disallows `require()` of ESM
  * Permits import assertions (but not in `nodenext`) ([Microsoft for Developers][6]).
* **Smarter inference for conditional return types** — catches mismatches when branches return differing types ([ssojet.com][7]).
* **`--erasableSyntaxOnly` flag** — lets Node.js run TS directly by stripping types-only syntax ([ssojet.com][7]).
* **`--libReplacement`** — allows substituting default lib files with custom ones ([ssojet.com][7]).
* https://github.com/microsoft/TypeScript/releases/tag/v5.8.2
* * https://devblogs.microsoft.com/typescript/announcing-typescript-5-8/

# 5.7 (Nov 2024)

* **`--target es2024` / `--lib es2024`** support:
  * `Object.groupBy`, `Map.groupBy`, `Promise.withResolvers`, etc.
* Smarter **`tsconfig.json` discovery** in editor mode
* https://github.com/microsoft/TypeScript/releases/tag/v5.7.2
* https://devblogs.microsoft.com/typescript/announcing-typescript-5-7/

# 5.6 (Sep 2024)

* **Disallowed nullish/truthy checks** — the compiler now errors when syntactically obvious always-truthy/nullish checks are used ([GitHub][1], [typescriptlang.org][2]).
* **New type `IteratorObject` & `--strictBuiltinIteratorReturn`** — more precise typing of built-in iterators ([typescriptlang.org][2]).
* **Support for arbitrary module identifiers** — import/export bindings with non-standard names (e.g. `export { banana as "" }`) ([typescriptdocs.com][3]).
* **`--noUncheckedSideEffectImports`** — flag to error on unresolved side‑effect imports ([typescriptdocs.com][3]).
* **New `--noCheck`** — allows skipping type checking for all input files ([AlternativeTo][4]).
* https://github.com/microsoft/TypeScript/releases/tag/v5.6.2
* https://devblogs.microsoft.com/typescript/announcing-typescript-5-6/

# 5.5 (Jun 2024)

* Improved **watch mode** & editor behavior (symlinks, `tsconfig`, file events)
* **Performance improvements**:

  * Faster control-flow analysis
  * Smaller output size
  * Speed-up in service APIs
* Better **discriminated union** handling via type caching

# 5.4 (Mar 2024)

* Checked **import attributes** using `ImportAttributes` interface
* Quick fix: **Add missing parameters**
* Enforced **deprecated features** from 5.0


# 5.3 (Nov 2023)

* Merged `typescript.js` + `tsserverlibrary.js` (20% size cut)
* Runtime check for **`super.field`** access
* Miscellaneous DOM lib & narrowing improvements

# 5.2 (Aug 2023)

* **Comma-aware** completions in objects
* **Inline variable** refactoring
* Smarter recursion checks in types
* Better **namespace emission** in declaration files

# 5.1 (Jun 2023)

* Easier **`undefined` return types** inference
* Decoupled **getter/setter types**
* Improved **JSX & React support**
* Editor support: **`@param` completions**, better performance

# 5.0 (Mar 2023)

* Support for **Decorators** (new standard syntax)
* **`--verbatimModuleSyntax`** flag: preserves import/export syntax exactly
* New **module resolution mode** for bundlers
* **Smaller compiler**, faster builds, requires ES2018+
