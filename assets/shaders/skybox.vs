#version 330

in vec3 vertexPosition;

uniform mat4 matModel;
uniform mat4 matView;
uniform mat4 matProjection;

out vec3 fragPosition;

void main()
{
    fragPosition = vertexPosition;
    mat4 matViewWithoutTranslation = mat4(mat3(matView));
    gl_Position = matProjection * matViewWithoutTranslation * matModel * vec4(vertexPosition, 1.0);
}
