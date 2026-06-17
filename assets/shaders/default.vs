#version 330

in vec3 vertexPosition;
in vec4 vertexColor;
in vec2 vertexTexCoord;
in vec3 vertexNormal;

uniform mat4 mvp;
uniform mat4 matModel;
uniform mat4 matNormal;
uniform vec3 viewPos;

out vec3 fragPosition;
out vec4 fragColor;
out vec2 fragTexCoord;
out vec3 fragNormal;

void main()
{
    fragPosition = vec3(matModel * vec4(vertexPosition, 1.0));
    fragTexCoord = vertexTexCoord;
    fragColor = vertexColor;
    fragNormal = normalize(vec3(matNormal * vec4(vertexNormal, 1.0)));
    gl_Position = mvp * vec4(vertexPosition, 1.0);
}

/*
#version 330

// input vertex attributes
in vec3 vertexPosition;

// input uniform values
uniform mat4 matProjection;
uniform mat4 matView;

// output vertex attributes (to fragment shader)
out vec3 fragPosition;

void main()
{
    // calculate fragment position based on model transformations
    fragPosition = vertexPosition;

    // remove translation from the view matrix
    mat4 rotView = mat4(mat3(matView));
    vec4 clipPos = matProjection * rotView * vec4(vertexPosition, 1.0);

    // calculate final vertex position
    gl_Position = clipPos;
}
*/
