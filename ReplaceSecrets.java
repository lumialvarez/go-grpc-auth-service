import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;

public class ReplaceSecrets {
	private static final String RUTA_ARCHIVOS = "src/cmd/devapi/config/envs/";
	private static final String ARCHIVO_TEMPLATE = "prod.env";
	private static final String ARCHIVO_PROPIEDADES = "prod.env";
	
	public static void main(String[] args) {
		try {
			if (args.length == 2) {
				String nombreVariable = args[0];
				String valorVariable = args[1];
				System.out.println("Variable: " + nombreVariable);

				String data = "";
				FileReader fileReader = null;

				File myObj = new File(RUTA_ARCHIVOS + ARCHIVO_PROPIEDADES);
				if (myObj.createNewFile()) {
					System.out.println("Cargando " + ARCHIVO_TEMPLATE);
					fileReader = new FileReader(new File(RUTA_ARCHIVOS + ARCHIVO_TEMPLATE).getAbsolutePath());
				} else {
					System.out.println("Cargando " + ARCHIVO_PROPIEDADES);
					fileReader = new FileReader(myObj.getAbsolutePath());
				}

				int i;
				while ((i = fileReader.read()) != -1) {
					data = data + ((char) i);
				}
				
				System.out.println("Reemplazando " + nombreVariable);
				data = data.replace("{" + nombreVariable + "}", valorVariable);
				
				System.out.println("Guardando " + ARCHIVO_PROPIEDADES);
				FileWriter myWriter = new FileWriter(myObj.getAbsolutePath());
				myWriter.write(data);
				myWriter.close();
			} else {
				System.out.println("Se requiere 2 valores: NombreVariable ValorVariable  (Args: " + args.length + ")");
			}

		} catch (Exception e) {
			System.out.println("An error occurred.");
			e.printStackTrace();
		}
	}
}