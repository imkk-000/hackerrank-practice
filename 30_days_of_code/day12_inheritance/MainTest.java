import java.lang.reflect.InvocationTargetException;

class MainTest {
    private static void TestClassStudentMethodNamecalculateInputFirstNameHeraldoLastNameMemelliID8135627Scores100And80ShouldReturnGradeO() {
        char expectedGrade = 'O';

        Student student = new Student("Heraldo", "Memelli", 8135627, new int[] { 100, 80 });
        char actualGrade = student.calculate();

        if (expectedGrade != actualGrade) {
            System.out.printf("expect %c but it got %c", expectedGrade, actualGrade);
        }
    }

    private static void TestClassStudentMethodNamecalculateInputFirstNameHeraldoLastNameMemelliID8135627Scores89ShouldReturnGradeE() {
        char expectedGrade = 'E';

        Student student = new Student("Heraldo", "Memelli", 8135627, new int[] { 89 });
        char actualGrade = student.calculate();

        if (expectedGrade != actualGrade) {
            System.out.printf("expect %c but it got %c", expectedGrade, actualGrade);
        }
    }

    private static void TestClassStudentMethodNamecalculateInputFirstNameHeraldoLastNameMemelliID8135627Scores79ShouldReturnGradeA() {
        char expectedGrade = 'A';

        Student student = new Student("Heraldo", "Memelli", 8135627, new int[] { 79 });
        char actualGrade = student.calculate();

        if (expectedGrade != actualGrade) {
            System.out.printf("expect %c but it got %c", expectedGrade, actualGrade);
        }
    }

    private static void TestClassStudentMethodNamecalculateInputFirstNameHeraldoLastNameMemelliID8135627Scores69ShouldReturnGradeP() {
        char expectedGrade = 'P';

        Student student = new Student("Heraldo", "Memelli", 8135627, new int[] { 69 });
        char actualGrade = student.calculate();

        if (expectedGrade != actualGrade) {
            System.out.printf("expect %c but it got %c", expectedGrade, actualGrade);
        }
    }

    private static void TestClassStudentMethodNamecalculateInputFirstNameHeraldoLastNameMemelliID8135627Scores79ShouldReturnGradeP() {
        char expectedGrade = 'D';

        Student student = new Student("Heraldo", "Memelli", 8135627, new int[] { 54 });
        char actualGrade = student.calculate();

        if (expectedGrade != actualGrade) {
            System.out.printf("expect %c but it got %c", expectedGrade, actualGrade);
        }
    }

    private static void TestClassStudentMethodNamecalculateInputFirstNameHeraldoLastNameMemelliID8135627Scores39ShouldReturnGradeT() {
        char expectedGrade = 'T';

        Student student = new Student("Heraldo", "Memelli", 8135627, new int[] { 39 });
        char actualGrade = student.calculate();

        if (expectedGrade != actualGrade) {
            System.out.printf("expect %c but it got %c", expectedGrade, actualGrade);
        }
    }

    public static void main(String[] args) {
        String testcaseName;
        for (var t : MainTest.class.getDeclaredMethods()) {
            testcaseName = t.getName();
            if (!testcaseName.startsWith("Test")) {
                continue;
            }
            try {
                System.out.printf("Run Test: %s\n", t.getName());
                t.invoke(null);
            } catch (IllegalAccessException ex) {
                ex.printStackTrace();
            } catch (InvocationTargetException ex) {
                ex.printStackTrace();
            }
        }
    }
}
