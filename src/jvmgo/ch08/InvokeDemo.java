package jvmgo.ch07;

public class InvokeDemo implements Runnable {
    public static int staticVar;
    public int instanceVar;
    public static void main(String[] args) {
        new InvokeDemo().test();
    }

    public void test() {
        InvokeDemo.startMethod();       // invokestatic
        InvokeDemo demo = new InvokeDemo(); // invokespecial
        demo.instanceMethod();          // invokespecial
        super.equals(null);              // invokespecial
        this.run();                     // invokevirtual
        ((Runnable) demo).run();         // invokeinterface
    }

    public static void startMethod() {}

    private void instanceMethod() {}


    @Override
    public void run() {

    }
}