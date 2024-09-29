def call(String name = "Caiman") {
    script {
        sh """
            echo Hi ${name}
        """
    }
}