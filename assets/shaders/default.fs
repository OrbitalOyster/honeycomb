#version 330

#define MAX_LIGHTS 2

in vec3 fragPosition;
in vec4 fragColor;
in vec2 fragTexCoord;
in vec3 fragNormal;

uniform vec3 ambientColor;

struct Light {
    int enabled;
    vec3 position;
    vec3 color;
    float radius;
};

uniform vec4 colDiffuse; /* Tint parameter in DrawModel() */
uniform vec3 viewPos;
uniform sampler2D texture0;
uniform Light lights[MAX_LIGHTS];

out vec4 finalColor;

void main()
{
    vec4 texelColor = texture(texture0, fragTexCoord);
    if (texelColor.a < 0.25) {
        discard;
    }

    vec3 norm = normalize(fragNormal);
    vec3 viewDir = normalize(viewPos - fragPosition);

    vec3 lighting = vec3(0.0, 0.0, 0.0);
    for (int i = 0; i < MAX_LIGHTS; i++) {
        /* Skip transparent parts */
        if (lights[i].enabled == 0)
            continue;
        vec3 lightDir = normalize(lights[i].position - fragPosition);
        /* Attenuation */
        float distance = length(lights[i].position - fragPosition);
        float attenuation = smoothstep(lights[i].radius, 0.0, distance);
        /* Diffuse */
        float diffuse = max(dot(norm, lightDir), 0.0);
        /* Specular */
        vec3 halfwayDir = normalize(lightDir + viewDir);
        float shininess = 32.0;
        float specularStrength = 1.0;
        float specular = specularStrength * pow(max(dot(norm, halfwayDir), 0.0), shininess);
        /* Done with lighting */
        lighting += (lights[i].color.rgb * diffuse + lights[i].color * specular) * attenuation;
    }

    /* Fresnel */
    float fresnelFactor = dot(norm, viewDir);
    float edgeFactor = clamp(1.0 - max(fresnelFactor, 0.0), 0.0, 1.0);
    float fresnel = pow(edgeFactor, 3.0);
    fresnel = 0.0;

    /* Fog */
    float fogDensity = 0.05;
    vec4 fogColor = vec4(0.4, 0.75, 1.0, 1.0);
    float distanceToCamera = length(viewPos - fragPosition);
    float fogFactor = 1.0 / exp((distanceToCamera * fogDensity) * (distanceToCamera * fogDensity));
    fogFactor = clamp(fogFactor, 0.0, 1.0);

    finalColor = texelColor * colDiffuse * vec4((ambientColor + lighting + fresnel), 1.0);

    finalColor = mix(fogColor, finalColor, fogFactor);
}
