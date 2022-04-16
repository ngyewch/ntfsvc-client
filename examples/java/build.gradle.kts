plugins {
    application
}

dependencies {
    implementation("com.github.ngyewch.ntfsvc:ntfsvc-client:1.0.0")
    implementation("org.apache.commons:commons-configuration2:2.7")
}

repositories {
    mavenCentral()
}

application {
    mainClass.set("Main")
}
