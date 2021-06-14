import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Arrays;
import org.wasmer.Instance;

public class Main {

    public static void main(String[] args) throws IOException {
        System.setProperty("os.arch", "amd64");
        String wasmFile = args[0];
        String functionName = args[1];
        Object[] functionArgs = Arrays.copyOfRange(args,2, args.length, Object[].class);

        Path wasmPath = Paths.get(wasmFile);
        byte[] wasmBytes = Files.readAllBytes(wasmPath);
        Instance instance = new Instance(wasmBytes);


        for (int i = 0; i < functionArgs.length; i++){
            try {
                functionArgs[i] = Integer.parseInt((String) functionArgs[i]);
            } catch (NumberFormatException ignored){}
        }

        Object[] results = instance.exports.getFunction(functionName).apply(functionArgs);

        System.out.println(Arrays.toString(results));
    }
}
