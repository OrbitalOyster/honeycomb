#version 330

in vec3 fragPosition;

uniform samplerCube skyboxMap;

out vec4 finalColor;

void main() {
    finalColor = texture(skyboxMap, fragPosition);
}
