A simple utility to execute a web assembly file via the JVM using Wasmer.

To Build:  
Install maven  
Install wasmer jar for your architecture to local repository:  
mvn install:install-file -Dfile=./wasmer-jni-<arch>-<platform>-0.3.0.jar  -DgroupId=org.wasmer -DartifactId=wasmer -Dversion=1.0.0 -Dpackaging=jar  

Then either  

Build and run the codebase:  
mvn clean install  
java -jar target/wasm-java-with-dependencies.jar <wasm file> <function name> <function args>  

Or   
Run the main method directly via an IDE such as intellij