# Videos

* `Introduction to shaders: Learn the basics`: https://www.youtube.com/watch?v=3mfvZ-mdtZQ

# Examples 

* [https://www.shadertoy.com/Protean_clouds](https://www.shadertoy.com/view/3l23Rh)

# Glossar
* [SDF](https://en.wikipedia.org/wiki/Signed_distance_function)
* [SPIR-V: Vulkan "bytecode" shading language](https://en.wikipedia.org/wiki/Standard_Portable_Intermediate_Representation)
    * can be compiled from
        * HLSL -> SPIR-V: [DirectXShaderCompiler SPIR-V CodeGen](https://github.com/microsoft/DirectXShaderCompiler)
        * GLSL -> SPIR-V: [GLSL Validator (SPIR-V Generator)](https://vulkan.lunarg.com/doc/view/1.0.39.1/linux/spirv_toolchain.html)
    * can [output HLSL and MSL](https://github.com/KhronosGroup/SPIRV-Cross)
        * Metal: GLSL -> SPIR-V -> MSL
        * HLSL: GLSL -> SPIR-V -> HLSL
