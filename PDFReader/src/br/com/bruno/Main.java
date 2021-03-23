package br.com.bruno;

import java.io.File;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import org.apache.pdfbox.pdmodel.PDDocument;
import org.apache.pdfbox.text.PDFTextStripper;

public class Main {
	public static void saveTextToFile(String text, String target) throws IOException {
		Path path = Paths.get(target.replace(".pdf", ".txt"));
		if (text.isBlank()) {
			System.out.println("Unable to get text from:" + target);
		}
		Files.writeString(path, text, StandardCharsets.UTF_8);
	}
	
	public static void main(String[] args) {
		if (args.length == 0) {
			System.out.println("usage: java -jar pdf-reader.jar pathToPDF");
			System.exit(1);
		}
		
		String path = args[0];
		
		File file = new File(path); 
		try {
			PDDocument document = PDDocument.load(file);
			
			saveTextToFile(new PDFTextStripper().getText(document), path.replace(".pdf", ".txt"));
			
			document.close();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
